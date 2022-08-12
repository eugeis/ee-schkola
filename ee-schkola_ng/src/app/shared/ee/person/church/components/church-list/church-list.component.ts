import {Component, OnInit} from '@angular/core';
import {ChurchDataService} from '../../services/church-data.service';
import {TableDataService} from '../../../../template/services/data.service';
import {Church} from '../../../../schkola/person/PersonApiBase';

@Component({
    selector: 'app-person-church-list',
    templateUrl: './church-list.component.html',
    styleUrls: ['./church-list.component.scss'],
    providers: [{provide: TableDataService, useClass: ChurchDataService}]
})

export class ChurchListComponent implements OnInit {

    church: Church = new Church();

    tableHeader: Array<String> = [];

    constructor(public churchDataService: ChurchDataService) { }

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader()
    }

    generateTableHeader() {
        return ['Actions', 'name', 'street', 'suite', 'city', 'code', 'country', 'first', 'last',
            'phone', 'email', 'cellphone', 'association', 'id'];
    }
}
