import {Component, OnInit} from '@angular/core';
import {ProfileViewService} from '../../services/profile-view.service';
import {ProfileDataService} from '../../services/profile-data.service';

@Component({
    selector: 'app-person-profile-list',
    templateUrl: './profile-list.component.html',
    styleUrls: ['./profile-list.component.scss'],
    providers: [ProfileViewService, ProfileDataService]
})

export class ProfileListComponent implements OnInit {

    constructor(public profileViewService: ProfileViewService, public profileDataService: ProfileDataService) { }

    ngOnInit(): void {
        this.profileViewService.init();
    }
}
