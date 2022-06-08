import {Component, OnInit} from '@angular/core';
import {ChurchViewService} from '../../services/church-view.service';
import {ChurchDataService} from '../../services/church-data.service';

@Component({
    selector: 'app-person-church',
    templateUrl: './church-view.component.html',
    styleUrls: ['./church-view.component.scss'],
    providers: [ChurchViewService, ChurchDataService]
})

export class ChurchViewComponent implements OnInit {

    constructor(public churchViewService: ChurchViewService, public churchDataService: ChurchDataService) {
    }

    ngOnInit(): void {

    }
}
