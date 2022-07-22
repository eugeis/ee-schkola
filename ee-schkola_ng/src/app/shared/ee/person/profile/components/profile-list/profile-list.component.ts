import {Component, OnInit} from '@angular/core';
import {ProfileDataService} from '../../services/profile-data.service';
import {TableDataService} from '../../../../template/services/data.service';
import {Profile} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-person-profile-list',
    templateUrl: './profile-list.component.html',
    styleUrls: ['./profile-list.component.scss'],
    providers: [{provide: TableDataService, useClass: ProfileDataService}]
})

export class ProfileListComponent implements OnInit {

    profile: Profile = new Profile();

    constructor(public profileDataService: ProfileDataService) { }

    ngOnInit(): void {

    }

    generateTableHeader() {
        const formArrayName = ['Actions'];
        Object.keys(this.profile).map((profileIndex) => {
            typeof this.profile[profileIndex] === 'object' ?
                this.profile[profileIndex] instanceof Date ?
                    formArrayName.push(profileIndex) :
                    Object.keys(this.profile[profileIndex]).forEach((element) => {
                        formArrayName.push(element);
                    }) : formArrayName.push(profileIndex);
        });
        return formArrayName;
    }
}
