import { HttpClient } from '@angular/common/http';
import { Component, OnInit, HostListener } from '@angular/core';
import {AbstractControl, ValidatorFn, Validators, FormControl, FormGroup } from '@angular/forms';
import { MatAutocomplete } from '@angular/material/autocomplete';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import { BackendLinkService } from '../services/backendservice/backend-link.service';

import jsonFile from './countries.json';

import { User } from './../dataModel/index.model';

interface Response {
  message: string;
}

@Component({
  selector: 'app-signup-page',
  templateUrl: './signup-page.component.html',
  styleUrls: ['./signup-page.component.scss'],
  host: {'class': 'default-layout'},
})
export class SignupPageComponent implements OnInit {
  constructor(private http: HttpClient, private backendLink: BackendLinkService) {
    const usrStr = localStorage.getItem('newUser');
    if (typeof usrStr === 'string') {
      this.newUser = JSON.parse(usrStr);
    } else {
      this.newUser = new User();
    }
  }

  private newUser: User;
  

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


  ngOnInit(): void {
    this.signupForm = new FormGroup({
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
      country : new FormControl('', {
        validators: [this.countryValidator(this.countries), Validators.required] 
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

    SignupPageComponent.fillForm(this.newUser, this.signupForm)
    
  }

  private _filter(value: string): string[] {
    const filterValue = value.toLowerCase();

    return this.countries.filter(country => country.toLowerCase().includes(filterValue));
  }

  onSubmit(form: FormGroup) {
    //this.newUser = this.formFieldsToObject(form);
    //console.log(this.newUser);
    this.newUser = SignupPageComponent.formGroupToUserObject(form);
    localStorage.setItem('newUser', JSON.stringify(this.newUser));
    return this.http
      .post<Response>(this.backendLink.getRegisterUrl(), this.newUser).subscribe({
        next: (data: Response) => {
          console.log(data);
          //TODO: show message and redirect
        },
        error: (err) => {
          console.error(err)
        }
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

