import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms'
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { MatSidenavModule} from '@angular/material/sidenav';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './sharepage/navbar/navbar.component';
import { FooterComponent } from './sharepage/footer/footer.component';
import { HomeComponent } from './pages/home/home.component';
import { AboutComponent } from './pages/about/about.component';
import { LoginComponent } from './pages/login/login.component';
import { GetStartedComponent } from './pages/get-started/get-started.component';import { SearchComponent } from './search/search.component';

import { AuthService } from 'app/auth.service';
import { AuthInterceptorService } from 'app/auth-interceptor.service';
import { CanActivateViaAuthGuard } from 'app/can-activate-via-auth.guard';
import { ChickenComponent } from './navbar-tabs/chicken/chicken.component';
import { BeefComponent } from './navbar-tabs/beef/beef.component';
import { SeafoodComponent } from './navbar-tabs/seafood/seafood.component';
import { PorkComponent } from './navbar-tabs/pork/pork.component';
//import { SeafoodComponent } from './navbar-tabs/seafood/seafood.component';
//import { BeefComponent } from './navbar-tabs/beef/beef.component';
import { userAccounts } from './user-accounts/user-accounts.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';


@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    FooterComponent,
    HomeComponent,
    AboutComponent,
    LoginComponent,
    GetStartedComponent,
    SearchComponent,
    ChickenComponent,
    BeefComponent,
    SeafoodComponent,
    PorkComponent,
    //SeafoodComponent,
    //BeefComponent,
   
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    BrowserAnimationsModule,
    MatSidenavModule,
    MatToolbarModule,
    MatIconModule,
    MatDividerModule
  ],
  providers: [
    AuthService,
    {
        provide: HTTP_INTERCEPTORS,
        useClass: AuthInterceptorService,
        multi: true
    },
    CanActivateViaAuthGuard],
  bootstrap: [AppComponent]
})
export class AppModule { }
