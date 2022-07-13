import {Component, OnInit} from '@angular/core';
import {ProfileViewService} from '../../services/profile-view.service';
import {TableDataService} from '../../../../template/services/data.service';
import {ProfileDataService} from '../../services/profile-data.service';
import {Gender} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-person-profile-view',
    templateUrl: './profile-view.component.html',
    styleUrls: ['./profile-view.component.scss'],
    providers: [ProfileViewService, {provide: TableDataService, useClass: ProfileDataService}]
})

export class ProfileViewComponent implements OnInit {

    genderEnum = this.profileViewService.loadEnumElement(Gender);

    constructor(public profileViewService: ProfileViewService) { }

    ngOnInit(): void {
        this.profileViewService.checkRoute();
    }
}
