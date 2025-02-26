import { EvaluationResult } from '../interfaces/evaluation-result';

export type SequenceEvent = {
  id: string;
  time: string;
  type: string;
};

export type SequenceStage = {
  image?: string;
  latestEvaluation?: EvaluationResult;
  latestEvent?: SequenceEvent;
  latestFailedEvent?: SequenceEvent;
  state: SequenceState;
  name: string;
};

export enum SequenceState {
  TRIGGERED = 'triggered',
  STARTED = 'started',
  FINISHED = 'finished',
  PAUSED = 'paused',
  TIMEDOUT = 'timedOut',
  ABORTED = 'aborted',
  SUCCEEDED = 'succeeded', //currently only for stages. It is actually like finished (it can still be failed)
  WAITING = 'waiting',
  UNKNOWN = '',
}

export enum SequenceStateControl {
  PAUSE = 'pause',
  ABORT = 'abort',
  RESUME = 'resume',
}

export class Sequence {
  name!: string;
  project!: string;
  service!: string;
  shkeptncontext!: string;
  stages!: SequenceStage[];
  state!: SequenceState;
  time!: string;
  problemTitle?: string;
}
