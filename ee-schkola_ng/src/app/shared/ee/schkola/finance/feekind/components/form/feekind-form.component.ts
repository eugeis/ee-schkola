import {FeeKind} from '@schkola/finance/FinanceApiBase'

import {Component, OnInit, Input} from '@angular/core';
import {FeeKindDataService} from '../../service/feekind-data.service';

@Component({
  selector: 'app-feekind-form',
  templateUrl: './feekind-form.component.html',
  styleUrls: ['./feekind-form.component.scss']
})

export class FeeKindFormComponent implements OnInit {


    @Input() feekind: FeeKind;

    constructor(public feekindDataService: FeeKindDataService) {}

    ngOnInit(): void {
        
    }
}




