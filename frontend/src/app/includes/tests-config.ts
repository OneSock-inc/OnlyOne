import { HttpClientModule } from "@angular/common/http";
import { HttpClientTestingModule } from "@angular/common/http/testing";
import { ReactiveFormsModule } from "@angular/forms";
import { MatAutocompleteModule } from "@angular/material/autocomplete";
import { SignupFormComponent } from "../forms/signup-form/signup-form.component";
import { LoaderComponent } from "../loader/loader.component";
import { LoaderDirective } from "../loader/loader.directive";
import { LoginPageComponent } from "../pages/login-page/login-page.component";
import { MesageBannerDirective } from "../message-banner/mesage-banner.directive";
import { MessageBannerComponent } from "../message-banner/message-banner.component";
import { services } from "../services";

export const testConfig = {
    imports: [HttpClientTestingModule, HttpClientModule, ReactiveFormsModule, MatAutocompleteModule],
    declarations: [
      SignupFormComponent,
      LoginPageComponent,
      MesageBannerDirective,
      MessageBannerComponent,
      LoaderComponent,
      LoaderDirective,
    ],
    providers: [services]
  }