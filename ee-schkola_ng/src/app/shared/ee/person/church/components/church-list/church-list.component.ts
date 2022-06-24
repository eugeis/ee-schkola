import {Component, OnInit} from '@angular/core';
import {ChurchViewService} from '../../services/church-view.service';
import {ChurchDataService} from '../../services/church-data.service';
import {TableDataService} from '../../../../template/services/data.service';

@Component({
    selector: 'app-person-church-list',
    templateUrl: './church-list.component.html',
    styleUrls: ['./church-list.component.scss'],
    providers: [ChurchViewService, {provide: TableDataService, useClass: ChurchDataService}]
})

export class ChurchListComponent implements OnInit {

    constructor(public churchViewService: ChurchViewService, public churchDataService: ChurchDataService) { }

    ngOnInit(): void {

    }
}
