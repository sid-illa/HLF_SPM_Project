import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginService } from '../login.service';

@Component({
  selector: 'login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
  providers: [LoginService]
})
export class LoginComponent implements OnInit {
  loginForm!: FormGroup;
  isSubmitted = false;
  constructor(private router: Router, private formBuilder: FormBuilder, private login: LoginService) { }

  ngOnInit() {
    this.loginForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
  }
  onSubmit() {
    this.login.loginUser(this.loginForm.value).subscribe((res: any) => {
      sessionStorage.setItem('token','token')
      sessionStorage.setItem('userid','39')
      alert("Login Success")
      this.router.navigateByUrl('/home')
    })
  }
}
