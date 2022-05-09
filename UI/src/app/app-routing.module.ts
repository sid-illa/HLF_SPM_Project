import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthgaurdService } from './authgaurd.service';
import { CertificatesComponent } from './certificates/certificates.component';
import { HomeComponent } from './home/home/home.component';
import { LayoutComponent } from './layout/layout.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';

const routes: Routes = [
  {
    path: '', component: LayoutComponent, canActivate: [AuthgaurdService],
    children: [
      { path: '', redirectTo: 'home',pathMatch:'full'},
      { path: 'home', component: HomeComponent },
      { path: 'uploadcertificates', component: CertificatesComponent },
    ]
  },
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
