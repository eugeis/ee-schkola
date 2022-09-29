import {Profile} from '@schkola/person/PersonApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {ProfileDataService} from '../../service/profile-data.service';

@Component({
  selector: 'app-profile-view',
  templateUrl: './profile-entity-view.component.html',
  styleUrls: ['./profile-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ProfileDataService}]
})

export class ProfileViewComponent implements OnInit {

    profile: Profile;

    constructor(public profileDataService: ProfileDataService) {}

    ngOnInit(): void {
        this.profile = this.profileDataService.getFirst();
        this.profileDataService.checkRoute(this.profile);
    }
}




