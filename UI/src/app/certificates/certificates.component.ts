import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CertificatesService } from './certificates.service';

@Component({
  selector: 'app-certificates',
  templateUrl: './certificates.component.html',
  styleUrls: ['./certificates.component.scss']
})
export class CertificatesComponent implements OnInit {

  certificatesForm!: FormGroup;
  isSubmitted = false;
  certificateInfo: any = [];
  cerficateInfo: any;
  constructor(private formBuilder: FormBuilder, private certificateService: CertificatesService) { }

  ngOnInit() {

    this.certificatesForm = this.formBuilder.group({
      filename: ['', Validators.required],
      filepath: ['', Validators.required]
    });
    this.certificateService.getCertificate().subscribe((res: any) => {
      this.certificateInfo = res.data;
      console.log(res)
    })
  }
  onSubmit() {
    console.log(this.certificatesForm)
    this.certificateService.uploadCertificate(this.certificatesForm.value).subscribe((res: any) => {
      // console.log(res)
      // // this.certificatesForm
      alert('certficate created')
      // this.router.navigateByUrl('')
    })
  }
  subString(info: String) {
    return info.substr(1, 100).concat(' ...')
  }
  viewCertificate(obj: any) {
    this.cerficateInfo = obj.fileinfo
  }
}
