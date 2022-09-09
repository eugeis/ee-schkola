import {Component, Input} from '@angular/core';
import {FinanceViewService} from '../../services/finance-view.service}';

@Component({
  selector: 'app-finance',
  templateUrl: './finance-view.component.html',
  styleUrls: ['./finance-view.component.scss'],
  providers: [FinanceViewService]
})

export class FinanceViewComponent {

    @Input() pageName = 'FinanceComponent';
       
    constructor(public financeViewService: FinanceViewService) {}

}



