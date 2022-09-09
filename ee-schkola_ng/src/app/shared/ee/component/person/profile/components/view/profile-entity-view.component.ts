import {Gender, Profile} from '../person/PersonApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {ProfileDataService} from '../../services/profile-data.service';

@Component({
  selector: 'app-profile-view',
  templateUrl: './profile-view.component.html',
  styleUrls: ['./profile-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ProfileDataService}]
})

export class ProfileViewComponent implements OnInit {

    genderEnum = this.profileDataService.loadEnumElement(Gender);
    profile: Profile;

    constructor(public profileDataService: ProfileDataService) {}

    ngOnInit(): void {
        this.profile = this.profileDataService.getFirst();
        this.profileDataService.checkRoute(this.profile);
    }
}




