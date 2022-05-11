import {Component, OnInit} from '@angular/core';
import {ButtonService} from '../../services/button.service';
import {TemplateService} from '../../services/template.service';

@Component({
    selector: 'app-button',
    templateUrl: './button.component.html',
    styleUrls: ['./button.component.scss'],
    providers: [ButtonService]
})

export class ButtonComponent implements OnInit {

    constructor(public templateService: TemplateService, public buttonService: ButtonService) { }

    ngOnInit(): void {
    }
}
