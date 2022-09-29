import {Fee, FeeKind} from '@schkola/finance/FinanceApiBase'
import {Profile} from '@schkola/person/PersonApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {FeeDataService} from '../../service/fee-data.service';

@Component({
  selector: 'app-fee-form',
  templateUrl: './fee-form.component.html',
  styleUrls: ['./fee-form.component.scss']
})

export class FeeFormComponent implements OnInit {


    @Input() fee: Fee;

    constructor(public feeDataService: FeeDataService) {}

    ngOnInit(): void {
        if (this.fee.student === undefined) {
            this.fee.student = new Profile();
        }
        if (this.fee.kind === undefined) {
            this.fee.kind = new FeeKind();
        }
    }
}




