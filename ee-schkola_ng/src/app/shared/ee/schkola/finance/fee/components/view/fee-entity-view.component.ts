import {Fee} from '@schkola/finance/FinanceApiBase'

import {Component, Inject, OnInit} from '@angular/core';
import {TableDataService} from '../../../../../template/services/data.service';
import {FeeDataService} from '../../service/fee-data.service';

@Component({
  selector: 'app-fee-view',
  templateUrl: './fee-entity-view.component.html',
  styleUrls: ['./fee-entity-view.component.scss'],
  providers: [{provide: TableDataService, useClass: FeeDataService}]
})

export class FeeViewComponent implements OnInit {


    fee: Fee;

    constructor(@Inject(FeeDataService) public feeDataService: FeeDataService) {}

    ngOnInit(): void {
        this.fee = this.feeDataService.getFirst();
        
        this.feeDataService.checkRoute(this.fee);
    }
}




