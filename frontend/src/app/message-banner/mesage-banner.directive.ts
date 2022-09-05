import { Directive, ViewContainerRef } from '@angular/core';
import { MessageBannerComponent } from './message-banner.component';

@Directive({
  selector: '[appMessageBanner]'
})
export class MessageBannerDirective {

  constructor(private vcref: ViewContainerRef) {

  }

  displayMessage(message: string){
    const elem = this.vcref.createComponent(
      MessageBannerComponent
    );
    elem.instance.message = message;
  }

  hideMessage() {
    this.vcref.clear();
  }

}
