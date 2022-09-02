import { Component, OnInit, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { User } from 'src/app/dataModel/user.model';
import { countryValidator, postalCodeValidator } from './../customValidators';
import jsonFile from './countries.json';
import { Observable } from 'rxjs/internal/Observable';
import {map, startWith} from 'rxjs/operators';
import { UserService } from 'src/app/services/userService/user-service.service';
import { MesageBannerDirective as MessageBannerDirective } from 'src/app/message-banner/mesage-banner.directive';
import { Router } from '@angular/router';

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

  // To display the list of countries
  countries: string[] = jsonFile.listOfCountries.map((country) => country.name);
  filteredCountries!: Observable<string[]>;

  @ViewChild(MessageBannerDirective, { static: true })
  messageBanner!: MessageBannerDirective;

  onSubmit(form: FormGroup): void {
    this.messageBanner.vcref.clear();
    console.log(SignupFormComponent.formGroupToUserObject(form));
    this.userService.registerNewUser(
      SignupFormComponent.formGroupToUserObject(form),
      (successMsg: any) => {
        console.log(successMsg);
        this.router.navigate(['/login']);
      }
      ,
      (errorMSg: any) => {
        this.messageBanner.displayMessage(errorMSg);
      }
    );
  }

  ngOnInit(): void {
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

    SignupFormComponent.fillForm(this.userService.getUser(), this.signupForm);

    this.filteredCountries = this.signupForm.controls[
      'country'
    ].valueChanges.pipe(
      startWith(''),
      map((value) => this._filter(value || ''))
    );
  }

  private _filter(value: string): string[] {
    const filterValue = value.toLowerCase();
    return this.countries.filter((country) =>
      country.toLowerCase().includes(filterValue)
    );
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

  private static fillForm(user: User, form: FormGroup): void {
    form.setValue({
      username: user.username,
      firstname: user.firstname,
      surname: user.surname,
      password: user.password,
      street: user.address.street,
      country: user.address.country,
      city: user.address.city,
      postalCode: user.address.postalCode,
    });
  }
  
}
