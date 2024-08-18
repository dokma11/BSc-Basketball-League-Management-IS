import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InterviewProposalCardComponent } from './interview-proposal-card.component';

describe('InterviewProposalCardComponent', () => {
  let component: InterviewProposalCardComponent;
  let fixture: ComponentFixture<InterviewProposalCardComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [InterviewProposalCardComponent]
    });
    fixture = TestBed.createComponent(InterviewProposalCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
