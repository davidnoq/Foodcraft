import { GetStartedComponent } from './pages/get-started/get-started.component';
import { AboutComponent } from './pages/about/about.component';
import { HomeComponent } from './pages/home/home.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SearchComponent } from './search/search.component';
import { LoginComponent } from './pages/login/login.component';
import { userAccounts } from './user-accounts/user-accounts.component';
import { CanActivateViaAuthGuard } from './can-activate-via-auth.guard';

const routes: Routes = [
  {path:'',component:HomeComponent},
  {path:'about',component:AboutComponent},
  {path:'get-started',component:GetStartedComponent},
  {path:'search',component:SearchComponent},
  {path:'login', component:LoginComponent},
  {path: 'profile', component:userAccounts, canActivate: [CanActivateViaAuthGuard]}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
