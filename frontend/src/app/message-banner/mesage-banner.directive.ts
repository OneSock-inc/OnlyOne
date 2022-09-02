import { Directive, ViewContainerRef } from '@angular/core';
import { MessageBannerComponent } from './message-banner.component';

@Directive({
  selector: '[appMessageBanner]'
})
export class MesageBannerDirective {

  constructor(public vcref: ViewContainerRef) {

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
