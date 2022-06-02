import {Component, OnInit} from '@angular/core';
import {ChurchViewService} from '../../services/church-view.service';

@Component({
    selector: 'app-person-profile',
    templateUrl: './church-view.component.html',
    styleUrls: ['./church-view.component.scss'],
    providers: [ChurchViewService]
})

export class ChurchViewComponent implements OnInit {

    constructor(public churchViewService: ChurchViewService) {
    }

    ngOnInit(): void {
        this.churchViewService.init();
    }
}
