import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'frontend';
  buttons = [
    {index:0, text: "login"},
    {index:1, text: "Add a lonely sock"},
    {index:2, text: "My socks"}
  ]
}
