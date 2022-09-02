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

  newUser: User = this.userService.getUser();

  // To display the list of countries
  countries: string[] = jsonFile.listOfCountries.map((country) => country.name);
  filteredCountries!: Observable<string[]>;

  @ViewChild(MessageBannerDirective, { static: true })
  messageBanner!: MessageBannerDirective;

  onSubmit(form: FormGroup): void {
    if (!form.valid) return
    this.messageBanner.hideMessage();
    this.userService.registerNewUser(
      SignupFormComponent.formGroupToUserObject(form),
      this.onSuccess,
      this.onError
    );
  }

  private onSuccess = (successMsg: any) => {
    console.log(successMsg);
    this.router.navigate(['/login']);
  }

  private onError = (errorMSg: any) => {
    this.messageBanner.displayMessage(errorMSg);
  }

  ngOnInit(): void {
    this.signupForm = new FormGroup({
      username: new FormControl( this.newUser.username, {
        validators: [Validators.required],
      }),
      password: new FormControl(this.newUser.password, {
        validators: [
          Validators.required,
          Validators.minLength(this.passwordMinLength),
        ],
      }),
      firstname: new FormControl(this.newUser.firstname, {
        validators: [Validators.required],
      }),
      surname: new FormControl(this.newUser.surname, {
        validators: [Validators.required],
      }),
      street: new FormControl(this.newUser.address.street, {
        validators: [Validators.required],
      }),
      country: new FormControl(this.newUser.address.country, {
        validators: [countryValidator(this.countries), Validators.required],
      }),
      postalCode: new FormControl(this.newUser.address.postalCode, {
        validators: [Validators.required, postalCodeValidator()],
      }),
      city: new FormControl(this.newUser.address.city, {
        validators: [Validators.required],
      }),
    });

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
  
}
