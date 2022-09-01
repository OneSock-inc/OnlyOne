import { HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, throwError } from 'rxjs';

/**
 * Service dedicated to http errors handling. All request pass throught this inteceptor.
 * See https://angular.io/api/common/http/HttpInterceptor
 */
@Injectable({
  providedIn: 'root'
})
export class HttpErrorService implements HttpInterceptor {
  constructor() { }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
      return next.handle(req).pipe(catchError(this.handleError))
  }

  private handleError(error: HttpErrorResponse) {
    let errMsg: string = '';
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      errMsg = 'A client-side error occured. Please verify your connection.';
    } else {
      errMsg = error.error?.message ? error.error.message : error.message;
    }

    // Return an observable with a user-facing error message.
    return throwError(() => new Error(errMsg));
  }
  
}
