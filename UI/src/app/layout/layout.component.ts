import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from '../login.service';

@Component({
  selector: 'app-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss']
})
export class LayoutComponent implements OnInit {

  constructor(private router: Router, private loginService: LoginService) { }
  userDetails: any;
  ngOnInit(): void {
    this.getCertificateDetails()
  }
  logout() {
    sessionStorage.clear();
    this.router.navigateByUrl('/login')
  }
  getCertificateDetails() {
    this.loginService.getUserDetails(sessionStorage.getItem('userid')).subscribe((res: any) => {
      console.log(res.data[0])
      this.userDetails = res.data[0] ? res.data[0] : {}
    })
  }
}
