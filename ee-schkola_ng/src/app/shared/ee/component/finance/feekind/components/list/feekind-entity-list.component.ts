import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {FeeKindDataService} from '../../services/feekind-data.service';

@Component({
  selector: 'app-feekind-list',
  templateUrl: './feekind-view.component.html',
  styleUrls: ['./feekind-view.component.scss'],
  providers: [{provide: TableDataService, useClass: FeeKindDataService}]
})

export class FeeKindViewComponent implements OnInit {

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




