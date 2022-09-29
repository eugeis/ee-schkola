import {FeeKind} from '@schkola/finance/FinanceApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {FeeKindDataService} from '../../service/feekind-data.service';

@Component({
  selector: 'app-feekind-list',
  templateUrl: './feekind-entity-list.component.html',
  styleUrls: ['./feekind-entity-list.component.scss'],
  providers: [{provide: TableDataService, useClass: FeeKindDataService}]
})

export class FeeKindListComponent implements OnInit {

    feekind: FeeKind = new FeeKind();

    tableHeader: Array<String> = [];

    constructor(public feekindDataService: FeeKindDataService) {}

    ngOnInit(): void {
        this.tableHeader = this.generateTableHeader();
    }

    generateTableHeader() {
        return ['Actions', 'name', 'amount', 'description', 'id'];
    }
}




