import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginService } from '../login.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss'],
  providers:[LoginService]
})
export class RegisterComponent implements OnInit {
  registerForm!: FormGroup;
  isSubmitted = false;
  constructor(private router: Router, private formBuilder: FormBuilder,private register:LoginService) { }

  ngOnInit() {
    this.registerForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
  }
  onSubmit() {
    console.log(this.registerForm)
    this.register.registerUser(this.registerForm.value).subscribe((res:any)=>{
      console.log(res)
      alert('user created')
      // this.router.navigateByUrl('')
    })
  }
}
