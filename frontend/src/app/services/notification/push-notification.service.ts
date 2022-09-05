import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { SwPush } from '@angular/service-worker';
import { BackendLinkService } from '../backendservice/backend-link.service';

@Injectable({
  providedIn: 'root'
})
export class PushNotificationService {
  readonly VAPID_PUBLIC_KEY : string = "BFGS8tken0rfjH_R04JzZb3eTWqkmAhVQJOvKH-HRN1sor_WEiijXIxahfgzIRr70V1MbB_lU4tLswZmtGmR3q4";
  
  constructor(private _swPush: SwPush,private _backend :BackendLinkService,private http: HttpClient) { 
    
  }
  requestSubscription = () => {
    Notification.requestPermission().then((permission) => {
        console.log("Notification perm asked.");
    });
    this._swPush.requestSubscription({
      serverPublicKey: this.VAPID_PUBLIC_KEY
    }).then(sub => {
      console.log("sub : %s", JSON.stringify(sub));
      this.http.post<any>(this._backend.getPushSubscriptionUrl(),sub).subscribe({
        next: response => {
          console.log("added this sub : %s", JSON.stringify(sub));
          console.log("Added subscription to backend");
          console.log(response);
        },
        error: (err) => {
          console.log("Added subscription to backend");
          console.log(err);
        }
      });
    }).catch((_) => console.log).catch(err => console.error("Could not subscribe to notifications", err));
  };
}
