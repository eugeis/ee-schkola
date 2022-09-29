import {Address, ChurchInfo, Contact, Education, Family, Gender, PersonName, Profile} from '@schkola/person/PersonApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {ProfileDataService} from '../../service/profile-data.service';

@Component({
  selector: 'app-profile-form',
  templateUrl: './profile-form.component.html',
  styleUrls: ['./profile-form.component.scss']
})

export class ProfileFormComponent implements OnInit {

    genderEnum = this.profileDataService.loadEnumElement(Gender);
    @Input() profile: Profile;

    constructor(public profileDataService: ProfileDataService) {}

    ngOnInit(): void {
        if (this.profile.name === undefined) {
            this.profile.name = new PersonName();
        }
        if (this.profile.address === undefined) {
            this.profile.address = new Address();
        }
        if (this.profile.contact === undefined) {
            this.profile.contact = new Contact();
        }
        if (this.profile.family === undefined) {
            this.profile.family = new Family();
        }
        if (this.profile.church === undefined) {
            this.profile.church = new ChurchInfo();
        }
        if (this.profile.education === undefined) {
            this.profile.education = new Education();
        }
    }
}




