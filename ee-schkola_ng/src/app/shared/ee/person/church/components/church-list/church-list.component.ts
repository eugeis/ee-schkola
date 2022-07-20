import {Component, OnInit} from '@angular/core';
import {ChurchDataService} from '../../services/church-data.service';
import {TableDataService} from '../../../../template/services/data.service';

@Component({
    selector: 'app-person-church-list',
    templateUrl: './church-list.component.html',
    styleUrls: ['./church-list.component.scss'],
    providers: [{provide: TableDataService, useClass: ChurchDataService}]
})

export class ChurchListComponent implements OnInit {

    constructor(public churchDataService: ChurchDataService) { }

    ngOnInit(): void {

    }
}
