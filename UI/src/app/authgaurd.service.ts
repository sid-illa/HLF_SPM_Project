import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthgaurdService {

  constructor(private router:Router) { }

  public canActivate(){
    if(sessionStorage.getItem('token')){
      return true;
    }
    else{
      this.router.navigateByUrl('/login')
      return false
    }
  }
}
