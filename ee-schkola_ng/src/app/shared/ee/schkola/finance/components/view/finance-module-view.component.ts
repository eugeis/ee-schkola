import {Component, Input} from '@angular/core';
import {FinanceViewService} from '../../service/finance-module-view.service';

@Component({
  selector: 'app-finance',
  templateUrl: './finance-module-view.component.html',
  styleUrls: ['./finance-module-view.component.scss'],
  providers: [FinanceViewService]
})

export class FinanceViewComponent {

    @Input() pageName = 'FinanceComponent';
       
    constructor(public financeViewService: FinanceViewService) {}

}



