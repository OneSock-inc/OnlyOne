import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { User } from 'src/app/dataModel/user.model';
import { countryValidator, postalCodeValidator } from './../customValidators';
import jsonFile from './countries.json';
import { Observable } from 'rxjs/internal/Observable';
import {catchError, map, startWith} from 'rxjs/operators';
import { UserService } from 'src/app/services/userService/user-service.service';
import { MessageBannerDirective } from 'src/app/message-banner/mesage-banner.directive';
import { Router } from '@angular/router';
import { of, throwError } from 'rxjs';

@Component({
  selector: 'app-signup-form',
  templateUrl: './signup-form.component.html',
  styleUrls: ['./signup-form.component.scss'],
})
export class SignupFormComponent implements OnInit {
  
  constructor(private userService: UserService, private router: Router) {}

  // Accessed in template
  signupForm!: FormGroup;
  hidePassword = true;
  passwordMinLength: number = 10;

  @Input() textButton?: string;
  @Input() isSignup?: boolean;
  @Input() submitBtnText!: string;

  user$!: Observable<User>;

  // To display the list of countries
  countries: string[] = jsonFile.listOfCountries.map((country) => country.name);
  filteredCountries!: Observable<string[]>;

  @ViewChild(MessageBannerDirective, { static: true })
  messageBanner!: MessageBannerDirective;

  ngOnInit(): void {
    this.isSignup = this.isSignup !== undefined;

    this.user$ = this.userService.getCurrentUser().pipe(
      map((data: User) => {
        this.fillForm(data);
        return data;
      })
    );

    this.createForm();

    this.filteredCountries = this.signupForm.controls[
      'country'
    ].valueChanges.pipe(
      startWith(''),
      map((value) => this._filter(value || ''))
    );
  }

  onSubmit(form: FormGroup): void {
    if (!form.valid) return;
    this.userService.registerNewUser(
      SignupFormComponent.formGroupToUserObject(form),
      this.onSuccess,
      this.onError
    );
  }

  private onSuccess = (successMsg: any) => {
    alert(successMsg.message)
    this.router.navigate(['/login']);
  };

  private onError = (errorMSg: any) => {
    alert(errorMSg);
  };

  
  onSubmitSave(form: FormGroup): void {
    if (!form.valid) return;
    this.user$ = this.userService
      .updateUser(SignupFormComponent.formGroupToUserObject(form))
      .pipe(
        map((d: User) => {
          alert(`User ${d.username} successfully saved !`);
          return d;
        }),
        catchError((err) => {
          alert(err.message);
          return of(err)
        })
      );
  }

  private _filter(value: string): string[] {
    const filterValue = value.toLowerCase();
    return this.countries.filter((country) =>
      country.toLowerCase().includes(filterValue)
    );
  }

  private createForm(): void {
    this.signupForm = new FormGroup({
      username: new FormControl('', {
        validators: [Validators.required],
      }),
      password: new FormControl('', {
        validators: [
          Validators.required,
          Validators.minLength(this.passwordMinLength),
        ],
      }),
      firstname: new FormControl('', {
        validators: [Validators.required],
      }),
      surname: new FormControl('', {
        validators: [Validators.required],
      }),
      street: new FormControl('', {
        validators: [Validators.required],
      }),
      country: new FormControl('', {
        validators: [countryValidator(this.countries), Validators.required],
      }),
      postalCode: new FormControl('', {
        validators: [Validators.required, postalCodeValidator()],
      }),
      city: new FormControl('', {
        validators: [Validators.required],
      }),
    });
  }

  private static formGroupToUserObject(form: FormGroup): User {
    const value = form.value;
    return {
      username: value.username,
      firstname: value.firstname,
      surname: value.surname,
      password: value.password,
      address: {
        street: value.street,
        country: value.country,
        city: value.city,
        postalCode: value.postalCode,
      },
    };
  }

  private fillForm(user: User): void {
    this.signupForm.setValue({
      username: user.username,
      firstname: user.firstname,
      surname: user.surname,
      password: '',
      street: user.address.street,
      country: user.address.country,
      city: user.address.city,
      postalCode: user.address.postalCode,
    });
  }
}
