import { Directive, ViewContainerRef } from '@angular/core';

@Directive({
  selector: '[appLoader]'
})
export class LoaderDirective {

  constructor(public viewContainerRef: ViewContainerRef) { }

}
