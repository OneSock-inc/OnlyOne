import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { User } from 'src/app/dataModel/user.model';
import { countryValidator, postalCodeValidator } from './../customValidators';
import jsonFile from './countries.json';
import { MatAutocomplete } from '@angular/material/autocomplete';
import { Observable } from 'rxjs/internal/Observable';
import {map, startWith} from 'rxjs/operators';



@Component({
  selector: 'app-signup-form',
  templateUrl: './signup-form.component.html',
  styleUrls: ['./signup-form.component.scss'],
})
export class SignupFormComponent implements OnInit {
  constructor() { }

  private newUser!: User;
  
  // Accessed in template
  signupForm!: FormGroup;
  hidePassword = true;
  passwordMinLength: number = 10;
  
  // To display the list of countries
  countries: string[]= jsonFile.listOfCountries.map(country => country.name);
  filteredCountries!: Observable<string[]>;

  onSubmit(form: FormGroup): User {
    this.newUser = SignupFormComponent.formGroupToUserObject(form);
    console.log(this.newUser);
    return this.newUser;
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

    this.filteredCountries = this.signupForm.controls['country'].valueChanges.pipe(
      startWith(''),
      map(value => this._filter(value || '')),
    );

  }
  
  private _filter(value: string): string[] {
    const filterValue = value.toLowerCase();
    return this.countries.filter(country => country.toLowerCase().includes(filterValue));
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
        postalCode: value.postalCode 
      }
    }
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
      postalCode: user.address.postalCode
    });
  }
  
}
