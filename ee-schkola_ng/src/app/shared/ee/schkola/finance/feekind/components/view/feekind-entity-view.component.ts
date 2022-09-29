import {FeeKind} from '@schkola/finance/FinanceApiBase'

import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {FeeKindDataService} from '../../service/feekind-data.service';

@Component({
  selector: 'app-feekind-view',
  templateUrl: './feekind-entity-view.component.html',
  styleUrls: ['./feekind-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: FeeKindDataService}]
})

export class FeeKindViewComponent implements OnInit {

    feekind: FeeKind;

    constructor(public feekindDataService: FeeKindDataService) {}

    ngOnInit(): void {
        this.feekind = this.feekindDataService.getFirst();
        this.feekindDataService.checkRoute(this.feekind);
    }
}




