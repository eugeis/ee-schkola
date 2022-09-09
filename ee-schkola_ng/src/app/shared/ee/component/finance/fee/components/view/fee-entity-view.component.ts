import {Component, OnInit} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';
import {FeeDataService} from '../../services/fee-data.service';

@Component({
  selector: 'app-fee-view',
  templateUrl: './fee-view.component.html',
  styleUrls: ['./fee-view.component.scss'],
  providers: [{provide: TableDataService, useClass: FeeDataService}]
})

export class FeeViewComponent implements OnInit {


    fee: Fee;

    constructor(public feeDataService: FeeDataService) {}

    ngOnInit(): void {
        this.fee = this.feeDataService.getFirst();
        this.feeDataService.checkRoute(this.fee);
    }
}




