import {Component, Input, OnInit} from '@angular/core';
import {ButtonService} from '../../services/button.service';
import {FormService} from '../../services/form.service';

@Component({
    selector: 'app-button',
    templateUrl: './button.component.html',
    styleUrls: ['./button.component.scss'],
    providers: [ButtonService]
})

export class ButtonComponent implements OnInit {

    @Input() usedService: FormService;
    @Input() isEdit: boolean;
    @Input() itemIndex: number;

    constructor(public buttonService: ButtonService) { }

    ngOnInit(): void {
    }
}
