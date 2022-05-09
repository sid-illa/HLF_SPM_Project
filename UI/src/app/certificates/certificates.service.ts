import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CertificatesService {
  constructor(private httpClient: HttpClient) { }

  uploadCertificate(data: any) {
    return this.httpClient.post('http://localhost:3000/extractfile', data);
  }
  //

  getCertificate() {
    return this.httpClient.get('http://localhost:3000/certicatesdata');
  }
}
