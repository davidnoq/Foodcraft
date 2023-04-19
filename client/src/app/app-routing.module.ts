import { ChickenComponent } from './navbar-tabs/chicken/chicken.component';
import { BeefComponent } from './navbar-tabs/beef/beef.component';
import { SeafoodComponent } from './navbar-tabs/seafood/seafood.component';
import { PorkComponent } from './navbar-tabs/pork/pork.component';
import { GetStartedComponent } from './pages/get-started/get-started.component';
import { AboutComponent } from './pages/about/about.component';
import { HomeComponent } from './pages/home/home.component';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SearchComponent } from './search/search.component';
import { LoginComponent } from './pages/login/login.component';
import { userAccounts } from './user-accounts/user-accounts.component';
import { CanActivateViaAuthGuard } from './can-activate-via-auth.guard';
import { FeaturedComponent } from './pages/featured/featured.component';
import { UsersearchComponent } from './navbar-tabs/usersearch/usersearch.component';


const routes: Routes = [
  {path:'',component:HomeComponent},
  {path:'about',component:AboutComponent},
  {path:'get-started',component:GetStartedComponent},
  {path:'search',component:SearchComponent},
  {path:'login', component:LoginComponent},
  {path:'chicken', component:ChickenComponent }, 
  {path:'beef', component:BeefComponent }, 
  {path:'salmon', component:SeafoodComponent }, 
  {path:'pork', component:PorkComponent }, 
  {path:'featured', component:FeaturedComponent }, 
  {path:'usersearch/:query', component:UsersearchComponent }, 
  {path: 'profile', component:userAccounts, canActivate: [CanActivateViaAuthGuard]}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
