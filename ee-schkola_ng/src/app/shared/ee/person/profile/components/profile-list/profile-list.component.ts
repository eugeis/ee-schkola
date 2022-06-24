import {Component, OnInit} from '@angular/core';
import {ProfileViewService} from '../../services/profile-view.service';
import {ProfileDataService} from '../../services/profile-data.service';
import {TableDataService} from '../../../../template/services/data.service';

@Component({
    selector: 'app-person-profile-list',
    templateUrl: './profile-list.component.html',
    styleUrls: ['./profile-list.component.scss'],
    providers: [ProfileViewService, {provide: TableDataService, useClass: ProfileDataService}]
})

export class ProfileListComponent implements OnInit {

    constructor(public profileViewService: ProfileViewService, public profileDataService: TableDataService) { }

    ngOnInit(): void {

    }
}
