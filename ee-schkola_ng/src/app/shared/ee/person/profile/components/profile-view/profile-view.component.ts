import {Component, OnInit} from '@angular/core';
import {ProfileViewService} from '../../services/profile-view.service';
import {ProfileDataService} from '../../services/profile-data.service';

@Component({
    selector: 'app-person-profile-view',
    templateUrl: './profile-view.component.html',
    styleUrls: ['./profile-view.component.scss'],
    providers: [ProfileViewService, ProfileDataService]
})

export class ProfileViewComponent implements OnInit {

    constructor(public profileViewService: ProfileViewService, public profileDataService: ProfileDataService) { }

    ngOnInit(): void {

    }
}
