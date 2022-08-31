import { Component, OnInit, HostListener } from '@angular/core';
import {AbstractControl, ValidatorFn, Validators, FormControl, FormGroup } from '@angular/forms';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';

import jsonFile from './countries.json';


@Component({
  selector: 'app-signup-page',
  templateUrl: './signup-page.component.html',
  styleUrls: ['./signup-page.component.scss'],
  host: {'class': 'default-layout'}
})
export class SignupPageComponent implements OnInit {

  username!: string;
  password!: string;
  firstname!: string;
  surname!: string;
  street!: string;
  country!: string;
  postalCode!: string;

  //// Validators

  // to detect if the postal code is a number
  postalCodeValidator(): ValidatorFn {
    return (control: AbstractControl): { [key: string]: any } | null => {
      const re = /^[0-9]*$/;
      if (re.test(control.value)) {
        return null  /* valid option selected */
      }
      return { 'invalidPostalCode': { value: control.value } }
    }
  }

  // to detect if the country is valid
  countryValidator(validOptions: Array<string>): ValidatorFn {
    return (control: AbstractControl): { [key: string]: any } | null => {
      if (validOptions.indexOf(control.value) !== -1) {
        return null  /* valid option selected */
      }
      return { 'invalidAutocompleteString': { value: control.value } }
    }
  }

  displayArrow: boolean = true;
  hidePassword = true;

  passwordMinLength: number = 10;

  // display down arrow if the user has not scrolled to the bottom of the page
  @HostListener('window:scroll', ['$event'])
  onScroll(event: Event): void {
    if (window.pageYOffset >= (document.documentElement.scrollHeight - document.documentElement.clientHeight)) {
      this.displayArrow = false;
    }
    else {
      this.displayArrow = true;
    }
  }


  // To display the list of countries
  countries: string[]= jsonFile.listOfCountries.map(country => country.name);
  filteredCountries!: Observable<string[]>;
  

  signupForm!: FormGroup;

  constructor() { }

  ngOnInit(): void {
    this.signupForm = new FormGroup({
      country : new FormControl('', {
        validators: [this.countryValidator(this.countries), Validators.required] 
      }),
      username: new FormControl('', {
        validators: [Validators.required]
      }),
      password: new FormControl('', {
        validators: [Validators.required, Validators.minLength(this.passwordMinLength)]
      }),
      firstname: new FormControl('', {
        validators: [Validators.required]
      }),
      surname: new FormControl('', {
        validators: [Validators.required]
      }),
      street: new FormControl('', {
        validators: [Validators.required]
      }),
      postalCode: new FormControl('', {
        validators: [Validators.required, this.postalCodeValidator()]
      }),
      city: new FormControl('', {
        validators: [Validators.required]
      })
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

  onSubmit(form: any): void {
    alert("Account created successfully");
    // send to api
    //form.countryValidator
    //form.username
    //form.password

  }

}

