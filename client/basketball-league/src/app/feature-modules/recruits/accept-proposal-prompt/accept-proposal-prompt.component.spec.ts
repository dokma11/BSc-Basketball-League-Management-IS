import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AcceptProposalPromptComponent } from './accept-proposal-prompt.component';

describe('AcceptProposalPromptComponent', () => {
  let component: AcceptProposalPromptComponent;
  let fixture: ComponentFixture<AcceptProposalPromptComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AcceptProposalPromptComponent]
    });
    fixture = TestBed.createComponent(AcceptProposalPromptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
