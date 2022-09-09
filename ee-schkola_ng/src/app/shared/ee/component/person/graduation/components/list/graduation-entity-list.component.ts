import {Graduation} from '../person/PersonApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {GraduationDataService} from '../../services/graduation-data.service';

@Component({
  selector: 'app-graduation-list',
  templateUrl: './graduation-view.component.html',
  styleUrls: ['./graduation-view.component.scss'],
  providers: [{provide: TableDataService, useClass: GraduationDataService}]
})

export class GraduationViewComponent implements OnInit {

    graduation: Graduation = new Graduation();

    tableHeader: Array<String> = [];

    constructor(public graduationDataService: GraduationDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'level', 'id'];
    }
}




