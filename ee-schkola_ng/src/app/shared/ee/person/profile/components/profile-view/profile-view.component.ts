import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {ProfileDataService} from '../../services/profile-data.service';
import {Address, Gender, PersonName, Profile} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-person-profile-view',
    templateUrl: './profile-view.component.html',
    styleUrls: ['./profile-view.component.scss'],
    providers: [{provide: TableDataService, useClass: ProfileDataService}]
})

export class ProfileViewComponent implements OnInit {

    genderEnum = this.profileDataService.loadEnumElement(Gender);

    profile: Profile;

    constructor(public profileDataService: ProfileDataService) { }

    ngOnInit(): void {
        this.profile = this.profileDataService.getFirst();
        this.profile.name = new PersonName();
        this.profile.address = new Address();
        this.profileDataService.checkRoute(this.profile);
    }
}
