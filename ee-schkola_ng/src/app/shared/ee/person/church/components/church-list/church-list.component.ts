import {Component, OnInit} from '@angular/core';
import {ChurchViewService} from '../../services/church-view.service';
import {ChurchDataService} from '../../services/church-data.service';

@Component({
    selector: 'app-person-church-list',
    templateUrl: './church-list.component.html',
    styleUrls: ['./church-list.component.scss'],
    providers: [ChurchViewService, ChurchDataService]
})

export class ChurchListComponent implements OnInit {

    constructor(public churchViewService: ChurchViewService, public churchDataService: ChurchDataService) { }

    ngOnInit(): void {
        this.churchViewService.init();
    }
}
