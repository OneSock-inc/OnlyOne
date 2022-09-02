import { ValidatorFn, AbstractControl } from "@angular/forms";

// to detect if the postal code is a number
export function postalCodeValidator(): ValidatorFn {
    return (control: AbstractControl): { [key: string]: any } | null => {
      const re = /^[0-9]*$/;
      if (re.test(control.value)) {
        return null  /* valid option selected */
      }
      return { 'invalidPostalCode': { value: control.value } }
    }
  }

  // to detect if the country is valid
export function countryValidator(validOptions: Array<string>): ValidatorFn {
    return (control: AbstractControl): { [key: string]: any } | null => {
      if (validOptions.indexOf(control.value) !== -1) {
        return null  /* valid option selected */
      }
      return { 'invalidAutocompleteString': { value: control.value } }
    }
  }