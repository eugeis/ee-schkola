import {Address, ChurchInfo, Contact, Education, Family, Gender, PersonName, Profile} from '@schkola/person/PersonApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {ProfileDataService} from '../../service/profile-data.service';

@Component({
  selector: 'app-profile-view',
  templateUrl: './profile-entity-view.component.html',
  styleUrls: ['./profile-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ProfileDataService}]
})

export class ProfileViewComponent implements OnInit {

    genderEnum = this.profileDataService.loadEnumElement(Gender);
    profile: Profile;

    constructor(@Inject(ProfileDataService) public profileDataService: ProfileDataService) {}

    ngOnInit(): void {
        this.profile = this.profileDataService.getFirst();
        this.profile.name = new PersonName();
        this.profile.address = new Address();
        this.profile.contact = new Contact();
        this.profile.family = new Family();
        this.profile.church = new ChurchInfo();
        this.profile.education = new Education();
        this.profileDataService.checkRoute(this.profile);
    }
}




