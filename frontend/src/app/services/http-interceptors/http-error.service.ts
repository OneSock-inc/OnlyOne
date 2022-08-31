import { HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HttpErrorService implements HttpInterceptor {
  constructor() { }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
      return next.handle(req).pipe(catchError(this.handleError))
  }

  private handleError(error: HttpErrorResponse) {
    let errMsg: string = 'Something bad happened; please try again later.';

    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      errMsg = 'A client-side error occured. Please verify your connection.';
    } else if (error.status === 401) {
      // Login failure
      errMsg = 'Login failed. Invalid username or password.'
    } else if (error.status >= 402 && error.status < 500) {
      errMsg = 'Unothorized action.'
    } else {
      errMsg = 'Server error. Please retry later.'
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error(errMsg));
  }
  
}
