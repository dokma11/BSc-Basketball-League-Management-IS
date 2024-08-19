import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeclineProposalPromptComponent } from './decline-proposal-prompt.component';

describe('DeclineProposalPromptComponent', () => {
  let component: DeclineProposalPromptComponent;
  let fixture: ComponentFixture<DeclineProposalPromptComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DeclineProposalPromptComponent]
    });
    fixture = TestBed.createComponent(DeclineProposalPromptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
