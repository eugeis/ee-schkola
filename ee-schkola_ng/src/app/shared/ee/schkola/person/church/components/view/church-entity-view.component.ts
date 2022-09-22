import {Address, Church, Contact, PersonName} from '@schkola/person/PersonApiBase'

import {Component, Inject, OnInit} from '@angular/core';
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

    constructor(@Inject(ChurchDataService) public churchDataService: ChurchDataService) {}

    ngOnInit(): void {
        this.church = this.churchDataService.getFirst();
        this.church.address = new Address();
        this.church.pastor = new PersonName();
        this.church.contact = new Contact();
        this.churchDataService.checkRoute(this.church);
    }
}




