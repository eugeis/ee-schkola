import {Church} from '@schkola/person/PersonApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {ChurchDataService} from '../../service/church-data.service';

@Component({
  selector: 'app-church-view',
  templateUrl: './church-entity-view.component.html',
  styleUrls: ['./church-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: ChurchDataService}]
})

export class ChurchViewComponent implements OnInit {

    church: Church;

    constructor(public churchDataService: ChurchDataService) {}

    ngOnInit(): void {
        this.church = this.churchDataService.getFirst();
        this.churchDataService.checkRoute(this.church);
    }
}




