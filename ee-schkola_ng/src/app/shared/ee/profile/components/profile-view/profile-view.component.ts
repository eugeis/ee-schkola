import {Component, OnInit} from '@angular/core';
import {ProfileViewService} from '../../services/profile-view.service';

@Component({
    selector: 'app-person-profile',
    templateUrl: './profile-view.component.html',
    styleUrls: ['./profile-view.component.scss'],
    providers: [ProfileViewService]
})

export class ProfileViewComponent implements OnInit {

    constructor(public profileViewService: ProfileViewService) { }

    ngOnInit(): void {
        this.profileViewService.init();
    }
}
