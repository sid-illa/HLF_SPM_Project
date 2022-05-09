import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private httpClient: HttpClient) { }

  loginUser(data: any) {
    return this.httpClient.post('http://localhost:3000/login', data);
  }

  registerUser(data: any) {
    return this.httpClient.post('http://localhost:3000/register', data);
  }

  getUserDetails(userId: any) {
    return this.httpClient.get(`http://localhost:3000/usercertificate/${userId}`);
  }
  //   login() {
  //     return this.httpClient.get('http://localhost:3000/login');
  //   }
}
