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

    constructor(public churchDataService: ChurchDataService) { }

    ngOnInit(): void {

    }

    generateTableHeader() {
        const formArrayName = ['Actions'];
        Object.keys(this.church).map((churchIndex) => {
            typeof this.church[churchIndex] === 'object' ?
                this.church[churchIndex] instanceof Date ?
                    formArrayName.push(churchIndex) :
                    Object.keys(this.church[churchIndex]).forEach((element) => {
                        formArrayName.push(element);
                    }) : formArrayName.push(churchIndex);
        })
        return formArrayName;
    }
}
